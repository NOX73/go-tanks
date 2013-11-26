function GameMap ( $scope ) {

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

}
