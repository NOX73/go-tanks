angular.module('app').controller('TankForm', ['$scope', '$tank', function ( $scope, $tank ) {

  //$scope.Id = 0
  //$scope.X = 0
  //$scope.Y = 0
  //$scope.LeftMotor = 0
  //$scope.RightMotor = 0
  //$scope.Direction = 0
  //$scope.GunDirection = 0

  $scope.$watch( function(){return $tank.Id}, function(val){ $scope.Id = $tank.Id } )

  $scope.$watch( function(){return $tank.Coords.X}, function(val){ $scope.X = $tank.Coords.X } )
  $scope.$watch( function(){return $tank.Coords.Y}, function(val){ $scope.Y = $tank.Coords.Y } )

  $scope.$watch( function(){return $tank.LeftMotor}, function(val){ $scope.LeftMotor = $tank.LeftMotor } )
  $scope.$watch( function(){return $tank.RightMotor}, function(val){ $scope.RightMotor = $tank.RightMotor } )

  $scope.$watch( function(){return $tank.Direction}, function(val){ $scope.Direction = $tank.Direction } )
  $scope.$watch( function(){return $tank.Gun.Direction}, function(val){ $scope.GunDirection = $tank.Gun.Direction } )


  $scope.$on("tank:message", function(event, message){
    $scope.setTank(message.Tank)
  })

  $scope.$on("world:message", function(event, message){
    tanks = message.Tanks

    i = _.findIndex(tanks, {Id: $tank.Id})
    if ( i >= 0 ) { $scope.setTank(tanks[i]) }

    $scope.$digest()
  })

  $scope.setTank = function ( t ) {
    $tank.Id = t.Id
    $tank.RightMotor = t.RightMotor
    $tank.LeftMotor = t.LeftMotor
    $tank.Direction = t.Direction
    $tank.Coords.X = t.Coords.X
    $tank.Coords.Y = t.Coords.Y
    $tank.Gun.Direction = t.Gun.Direction
  }

  $scope.apply = function () {}
}])
