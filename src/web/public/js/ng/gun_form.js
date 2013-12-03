angular.module('app').controller('GunForm', ['$scope', '$tank', function ( $scope, $tank ) {
  $scope.$watch( function(){ return $tank.gun.Direction }, function(val){ $scope.Direction = $tank.gun.Direction } )

  $scope.apply = function () {

    $tank.gun.Direction = $scope.Direction

    $scope.sendTankCommand({
      Gun: { Direction: $scope.Direction }
    })
  }


}])
