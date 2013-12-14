angular.module('app').controller('GameMap', ['$scope', '$tank', function ( $scope, $tank ) {
  var stage = new createjs.Stage("gameMap");

  var tanks = {}
  var bullets = {}

  var tankSpriteSheet = new createjs.SpriteSheet({
    images: ["/public/img/tanks.png"],
    frames: [ [ 13, 133, 132, 83, 0, 65, 41 ] ]
  });

  var gunSpriteSheet = new createjs.SpriteSheet({
    images: ["/public/img/tanks.png"],
    frames: [ [ 40, 237, 93, 45, 0, 24, 22 ] ]
  });

  function Tank( tank ){
    this.id = tank.Id
    this.container = new createjs.Container()

    this.tankSprite = new createjs.Sprite(tankSpriteSheet, 0)
    this.gunSprite = new createjs.Sprite(gunSpriteSheet, 0)

    this.container.addChild(this.tankSprite)
    this.container.addChild(this.gunSprite)

    this.container.scaleX = 0.15
    this.container.scaleY = 0.15

    this.update( tank )
  }

  Tank.prototype = {

    update: function ( tank ) {

      this.container.x = tank.Coords.X
      this.container.y = tank.Coords.Y
      this.container.rotation = tank.Direction

      this.gunSprite.rotation = tank.Gun.Direction
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
        stage.addChild(tanks[ tank.Id ].container);
      }else{
        tanks[ tank.Id ].update( tank )
      }
    });

    _.forIn( tanks,  function( tank ) {

      if( tanksIds.indexOf( tank.id ) == -1 ){
        stage.removeChild( tank.container )
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
    $event.preventDefault();

    if($tank.RightMotor + 0.1 < 1)
      $tank.RightMotor += 0.1
    if($tank.LeftMotor - 0.1 > -1)
      $tank.LeftMotor -= 0.1
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })
  }

  $scope.pressUp = function ($event) {
    $event.preventDefault();

    if($tank.LeftMotor + 0.1 > 1){return}
    if($tank.RightMotor + 0.1 > 1){return}
    $tank.LeftMotor += 0.1
    $tank.RightMotor += 0.1
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })
  }

  $scope.pressRight = function ($event) {
    $event.preventDefault();


    if($tank.LeftMotor + 0.1 < 1)
      $tank.LeftMotor += 0.1
    if($tank.RightMotor - 0.1 > -1)
      $tank.RightMotor -= 0.1
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })
  }

  $scope.pressDown = function ($event) {
    $event.preventDefault();

    if($tank.LeftMotor - 0.1 < -1){return}
    if($tank.RightMotor - 0.1 < -1){return}
    $tank.LeftMotor -= 0.1
    $tank.RightMotor -= 0.1
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })
  }

  $scope.pressSpace = function ($event) {
    $event.preventDefault();

    $tank.LeftMotor = 0
    $tank.RightMotor = 0
    $scope.sendTankCommand({ LeftMotor: $tank.LeftMotor, RightMotor: $tank.RightMotor })
  }

  $scope.pressEnter = function ($event) {
    $scope.sendTankCommand({ Fire: true })

    $event.preventDefault();
  }
}]);
