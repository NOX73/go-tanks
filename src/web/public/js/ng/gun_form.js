angular.module('app').controller('GunForm', ['$scope', function ( $scope ) {
  $scope.TurnAngle = 0

  $scope.apply = function () {

    $scope.sendTankCommand({
      Gun: { TurnAngle: $scope.TurnAngle }
    })

  }


}])
