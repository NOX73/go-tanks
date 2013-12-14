angular.module('app').controller('SessionForm', function ( $scope ) {

  $scope.WorldFrequency = parseInt(localStorage.worldFrequency) || 0

  $scope.resetSession = function () {
    $scope.closeSession()
    $scope.initSocket()
    $scope.messages.length = 0
  }

  $scope.closeSession = function () {
    $scope.closeSocket();
  }

  $scope.setFreqency = function () {
    localStorage.worldFrequency = parseInt($scope.WorldFrequency)
    $scope.sendFreqency($scope.WorldFrequency)
  }

});
