angular.module('app').controller('GunForm', ['$scope', function ( $scope ) {
  $scope.TurnAngle = 0

  $scope.apply = function () {

    console.log($scope.TurnAngle)

    $scope.sendTankCommand({
      Gun: { TurnAngle: $scope.TurnAngle }
    })

  }


}])
