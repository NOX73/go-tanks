angular.module('app').controller('WebSocks', function ( $scope ) {

  $scope.state = 'auth'
  $scope.messages = [ ]
  var socket;

  $scope.initSocket = function () {
    socket = new WebSocket("ws://" + window.location.host + "/ws")

    socket.onmessage = function ( event ) {
      message = JSON.parse(event.data)

      if( message["Message"] ){
        $scope.messages.push({ text: message["Message"] })
        if( $scope.messages.length > 5 ){ $scope.messages.shift() }
        $scope.$digest();
      }

      switch(message["Type"]){
        case "Wrold":
          $scope.$broadcast('world:message', message)
          break;

        case "Tank":
          $scope.$broadcast('tank:message', message)
          break;

        case "Pong":
          $scope.$broadcast('message:pong', message)
          break;
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
    //$scope.messageText = '';
  }

  $scope.sendTankCommand = function ( command ) {
    command.Type = "TankCommand"
    $scope.sendMessage(command)
  }

  $scope.sendClientCommand = function ( command ) {
    command.Type = "Client"
    $scope.sendMessage(command)
  }

  $scope.isAuth = function () { return $scope.state == 'auth' }
  $scope.isMessage = function () { return $scope.state == 'message' }
  $scope.isSession = function () { return $scope.state == 'session' }
  $scope.isControl = function () { return $scope.state == 'control' }
  $scope.isTank = function () { return $scope.state == 'tank' }

  $scope.setAuth = function () { $scope.state = 'auth' }
  $scope.setMessage = function () { $scope.state = 'message' }
  $scope.setSession = function () { $scope.state = 'session' }
  $scope.setControl = function () { $scope.state = 'control' }
  $scope.setTank = function () { $scope.state = 'tank' }

  $scope.$on('auth:success', function() {
    $scope.setControl()
    if (localStorage.worldFrequency)
      $scope.sendFreqency(localStorage.worldFrequency)
  } )

  $scope.sendFreqency = function (val) {
    $scope.sendClientCommand({ WorldFrequency: parseInt(val) })
  }
});
