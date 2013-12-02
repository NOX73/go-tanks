angular.module('app').controller('GameMap', ['$scope', '$tank', function ( $scope, $tank ) {

  var stage = new createjs.Stage("gameMap");

  var tanks = {}
  var bullets = {}

  function Tank( tank ){
    this.id = tank.Id
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

  function Bullet( bullet ){
    this.id = bullet.Id
    this.circle = new createjs.Shape();
    this.circle.graphics.beginFill("green").drawCircle(0, 0, 2);
    this.update( bullet )
  }

  Bullet.prototype = {

    update: function ( bullet ) {
      this.circle.x = bullet.Coords.X
      this.circle.y = bullet.Coords.Y
    }

  }

  $scope.$on( 'world:message', function( event, message ) {
    var tanksIds = []
    message.Tanks.forEach(function(tank){
      tanksIds.push( tank.Id )

      if( !tanks[ tank.Id ] ) {
        tanks[ tank.Id ] = new Tank( tank )
        stage.addChild(tanks[ tank.Id ].circle);
      }else{
        tanks[ tank.Id ].update( tank )
      }
    });

    _.forIn( tanks,  function( tank ) {

      if( tanksIds.indexOf( tank.id ) == -1 ){
        stage.removeChild( tank.circle )
        delete tanks[ tank.id ]
      }

    });

    var bulletsIds = []
    message.Bullets.forEach( function(bullet) {
      bulletsIds.push( bullet.Id )

      if( !bullets[bullet.Id] ) {
        bullets[ bullet.Id ] = new Bullet( bullet )
        stage.addChild(bullets[ bullet.Id ].circle);
      } else {
        bullets[ bullet.Id ].update( bullet )
      }
    });

    _.forIn( bullets,  function( bullet ) {

      if( bulletsIds.indexOf( bullet.id ) == -1 ){
        stage.removeChild( bullet.circle )
        delete bullets[ bullet.id ]
      }

    });

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

  $scope.pressEnter = function ($event) {
    $scope.sendTankCommand({ Fire: true })

    $event.preventDefault();
  }
}]);
