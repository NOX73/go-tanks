set :application, 'go_tanks'
set :repo_url, 'https://github.com/NOX73/go-tanks.git'

# ask :branch, proc { `git rev-parse --abbrev-ref HEAD`.chomp }

 set :deploy_to, '/u/apps/go_tanks'
 set :scm, :git

# set :format, :pretty
# set :log_level, :debug
# set :pty, true
set :stage, :production

# set :linked_files, %w{config/database.yml}
# set :linked_dirs, %w{bin log tmp/pids tmp/cache tmp/sockets vendor/bundle public/system}

# set :default_env, { path: "/opt/ruby/bin:$PATH" }
 set :keep_releases, 5

 set :go_tanks_pid, shared_path.join('go_tanks.pid')
 set :go_tanks_log, shared_path.join('go_tanks.log')
 set :go_tanks_main, release_path.join('main')
 # params for compile exec file
 set :go_build_os, 'linux'
 set :go_build_arch, 'amd64'

 logger.level = Logger::INFO

namespace :deploy do

  desc 'Restart application'
  task :restart do
    on roles(:app), in: :sequence, wait: 5 do
      # Your restart mechanism here, for example:
      # execute :touch, release_path.join('tmp/restart.txt')

      begin
        execute "kill -9 `cat #{fetch :go_tanks_pid}`"
      rescue Exception
      end

      execute "GOPATH=#{release_path} nohup #{fetch :go_tanks_main} > #{fetch :go_tanks_log} 2>&1 & echo $! > #{fetch :go_tanks_pid}"

    end
  end

  desc 'Compile main executable file for target os & architecture.'
  task :build do
    system({"GOPATH" =>Dir.pwd, "GOOS" => fetch(:go_build_os), "GOARCH" => fetch(:go_build_arch) }, "go build main.go")
  end

  desc 'Upload compiled main file to release.'
  task :upload_main do
    on roles(:app) do |host|
      upload!(File.join(Dir.pwd, 'main'), File.join(release_path, 'main'))
    end
  end

  after :restart, :clear_cache do
    on roles(:web), in: :groups, limit: 3, wait: 10 do
      # Here we can do anything such as:
      # within release_path do
      #   execute :rake, 'cache:clear'
      # end
    end
  end

  after :started, 'deploy:build'
  after :updated, 'deploy:upload_main'
  after :finishing, 'deploy:cleanup'

end
