angular.module('app').controller('SessionForm', function ( $scope ) {

  var pings = {}

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

  $scope.ping = function () {
    var id = 0
    pings[id] = Date.now()
    $scope.sendMessage({"Type":"Ping", "PingId":id})
  }

  $scope.$on("message:pong", function(event, message){
    $scope.pingValue = Date.now() - pings[message.PongId]
    delete pings[message.PongId]

    $scope.$digest()
  })

});
