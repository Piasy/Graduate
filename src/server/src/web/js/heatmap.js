/**
 * Created by piasy on 15/5/18.
 */

function updateHeatMap(players, heatmapOverlay) {
    players.forEach(function (p, i, arr) {
        if (p.selected) {
            var remark = JSON.parse(localStorage.getItem(p.history));
            if (remark != null && remark.position != null) {
                var points =[];

                remark.position.forEach(function (pp, ii, aaa) {
                    if (pp.latitude >= 0 && pp.longitude >= 0) {
                        points.push({"lng": pp.longitude, "lat": pp.latitude});
                    }
                });

                heatmapOverlay.setDataSet({data:points, max:10});
            }

        }
    });
}