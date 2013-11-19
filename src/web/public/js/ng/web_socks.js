var myApp = angular.module('myApp',[]);

myApp.controller('WebSocks', ['$scope', function ($scope, $ws) {

  $scope.messages = [ ]
  var socket = new WebSocket("ws://localhost:9000/ws"); 

  socket.onmessage = function ( event ) {
    $scope.messages.push({ text: event.data })
    if( $scope.messages.length > 5 ){ $scope.messages.shift() }
    $scope.$digest();
  }

  $scope.sendMessage = function () {
    socket.send( $scope.messageText );
    $scope.messageText = '';
  }

}])
