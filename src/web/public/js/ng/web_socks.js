angular.module('app').controller('WebSocks', function ( $scope ) {

  $scope.state = 'auth'
  $scope.messages = [ ]
  var socket;

  $scope.initSocket = function () {
    socket = new WebSocket("ws://localhost:9000/ws"); 

    socket.onmessage = function ( event ) {
      message = JSON.parse(event.data)

      if( message["Message"] ){
        $scope.messages.push({ text: message["Message"] })
        if( $scope.messages.length > 5 ){ $scope.messages.shift() }
        $scope.$digest();
      }

      if( message["Type"] == "World" ) {
        $scope.$broadcast('world:message', message)
      }
    }
  }

  $scope.closeSocket = function () {
    if( socket.readyState == 1 ){ socket.close(); }
  }

  $scope.initSocket()

  $scope.sendMessage = function ( message ) {
    if( !message ) { message = $scope.messageText } else { message = JSON.stringify( message ) }

    socket.send( message );
    $scope.messageText = '';
  }

  $scope.isAuth = function () { return $scope.state == 'auth' }
  $scope.isMessage = function () { return $scope.state == 'message' }
  $scope.isSession = function () { return $scope.state == 'session' }
  $scope.isControl = function () { return $scope.state == 'control' }

  $scope.setAuth = function () { $scope.state = 'auth' }
  $scope.setMessage = function () { $scope.state = 'message' }
  $scope.setSession = function () { $scope.state = 'session' }
  $scope.setControl = function () { $scope.state = 'control' }

  $scope.$on('auth:success', $scope.setControl )
});
