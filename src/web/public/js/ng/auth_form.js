function AuthForm($scope) {

  $scope.login = ''
  $scope.password = ''

  $scope.sendAuth = function () {
    $scope.sendMessage( JSON.stringify({ Type:"Auth", Login: $scope.login, Password: $scope.password }) )
  }

}
