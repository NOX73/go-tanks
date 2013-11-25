function WebSocks($scope) {

  $scope.state = 'auth'
  $scope.messages = [ ]
  var socket;

  function initSocket () {
    socket = new WebSocket("ws://localhost:9000/ws"); 

    socket.onmessage = function ( event ) {
      message = JSON.parse(event.data)

      if( message["Message"] ){
        $scope.messages.push({ text: message["Message"] })
        if( $scope.messages.length > 5 ){ $scope.messages.shift() }
        $scope.$digest();
      }
    }
  }

  function reinitSocket () {
    socket.close();
    initSocket();
  }

  initSocket()

  $scope.sendMessage = function ( message ) {
    if( !message ) { message = $scope.messageText }

    socket.send( message );
    $scope.messageText = '';
  }

  $scope.isAuth = function () { return $scope.state == 'auth' }
  $scope.isMessage = function () { return $scope.state == 'message' }
  $scope.isSession = function () { return $scope.state == 'session' }

  $scope.setAuth = function () { $scope.state = 'auth' }
  $scope.setMessage = function () { $scope.state = 'message' }
  $scope.setSession = function () { $scope.state = 'session' }

  $scope.$on('hello', function(){ $scope.state = "hello"; $scope.$digest() })

  $scope.resetSession = function () {
    reinitSocket()
    $scope.messages.length = 0
  }
}
