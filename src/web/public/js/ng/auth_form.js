function AuthForm($scope) {

  $scope.login = ''
  $scope.password = ''

  if(localStorage.login) {$scope.login = localStorage.login}
  if(localStorage.password) {$scope.password = localStorage.password}

  $scope.sendAuth = function () {
    localStorage.login = $scope.login
    localStorage.password = $scope.password

    $scope.sendMessage( { Type:"Auth", Login: $scope.login, Password: $scope.password } )

    $scope.$emit('auth:success')
  }

  if($scope.login != '' && $scope.password != ''){ setTimeout( $scope.sendAuth, 1000 ) }

}
