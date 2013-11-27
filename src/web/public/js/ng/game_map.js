angular.module('app').controller('GameMap', ['$scope', '$tank', function ( $scope, $tank ) {

  var stage = new createjs.Stage("gameMap");
  var tanks = {}

  function Tank( tank ){
    this.circle = new createjs.Shape();
    this.circle.graphics.beginFill("red").drawCircle(0, 0, 10);
    this.update( tank )
  }

  Tank.prototype = {

    update: function ( tank ) {
      this.circle.x = tank.Coords.X
      this.circle.y = tank.Coords.Y
    }

  }

  $scope.$on( 'world:message', function( event, message ) {
    message.Tanks.forEach(function(tank){
      if( !tanks[ tank.Id ] ) {

        tanks[ tank.Id ] = new Tank( tank )
        stage.addChild(tanks[ tank.Id ].circle);

      }else{

        tanks[ tank.Id ].update( tank )

      }
    })

    stage.update()
  })

  $scope.pressLeft = function ($event) {
    $tank.LeftMotor = 0.4
    $tank.RightMotor = 0.5
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })

    $event.preventDefault();
  }

  $scope.pressUp = function ($event) {
    $tank.LeftMotor = 1
    $tank.RightMotor = 1
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })

    $event.preventDefault();
  }

  $scope.pressRight = function ($event) {
    $tank.LeftMotor = 0.5
    $tank.RightMotor = 0.4
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })

    $event.preventDefault();
  }

  $scope.pressDown = function ($event) {
    $tank.LeftMotor = 0
    $tank.RightMotor = 0
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })

    $event.preventDefault();
  }
}]);
