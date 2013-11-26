angular.module('app').controller('SessionForm', function ( $scope ) {

  $scope.resetSession = function () {
    $scope.closeSession()
    $scope.initSocket()
    $scope.messages.length = 0
  }

  $scope.closeSession = function () {
    $scope.closeSocket();
  }


});
