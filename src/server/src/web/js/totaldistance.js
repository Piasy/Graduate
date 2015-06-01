/**
 * Created by piasy on 15/5/18.
 */

function createTotalDistanceChart(players) {
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });

    var player_names = [];
    var s = {};
    s.showInLegend = false;
    s.data = [];
    players.forEach(function (p, i, arr) {
        player_names.push(p.name);
        s.data.push(0);
    });
    var series_data = [];
    series_data.push(s);

    var series = [];
    $('#chart_container').highcharts({
        chart: {
            type: 'column',
            animation: Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
                load: function() {
                    series = this.series;
                }
            }
        },
        title: {
            text: '跑动总距离'
        },
        xAxis: {
            categories: player_names
        },
        yAxis: {
            min: 0,
            title: {
                text: '距离（米）'
            }
        },
        tooltip: {
            pointFormat: '<span><b>{point.y:.1f}</b> 米</span>',
            shared: true,
            useHTML: true
        },
        plotOptions: {
            column: {
                pointPadding: 0.2,
                borderWidth: 0
            }
        },
        exporting: {
            enabled: false
        },
        credits: {
            enabled: false
        },
        series: series_data
    });
    return series;
}

function updateTotalDistanceChartFirstTime(series, players) {
    var data = [];
    players.forEach(function (p, i, arr) {
        /*if (p.selected) {
         series[i].show();
         } else {
         series[i].hide();
         }*/
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.curdistance != undefined) {
            data.push(remark.curdistance);
        }
    });
    series[0].setData(data);
}

function updateTotalDistanceChart(series, index, remark) {
    /*if (player.selected) {
     series.show();
     } else {
     series.hide();
     }*/
    var points = series.points;
    var ys = [];
    points.forEach(function (p, i, arr) {
        ys.push(p.y);
    });
    if (remark != null && remark.curdistance != undefined && ys.length > index && ys[index] < remark.curdistance) {
        ys[index] = remark.curdistance;
        series.setData(ys);
    }
}
