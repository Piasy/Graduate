/**
 * Created by piasy on 15/5/18.
 */

function createDistWithSpeedChart(players) {
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });

    var player_names = [];
    players.forEach(function (p, i, arr) {
        player_names.push(p.name);
    });

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
            text: '不同速度跑动距离统计'
        },
        xAxis: {
            categories: player_names
        },
        yAxis: {
            min: 0,
            title: {
                text: '距离（米）'
            },
            stackLabels: {
                enabled: true,
                style: {
                    fontWeight: 'bold',
                    color: (Highcharts.theme && Highcharts.theme.textColor) || 'gray'
                }
            }
        },
        tooltip: {
            formatter: function() {
                return '<b>'+ this.x +'</b><br/>'+
                    this.series.name +': '+ this.y +'<br/>'+
                    '总距离：'+ this.point.stackTotal;
            }
        },
        plotOptions: {
            column: {
                stacking: 'normal',
                dataLabels: {
                    enabled: true,
                    color: (Highcharts.theme && Highcharts.theme.dataLabelsColor) || 'white'
                }
            }
        },
        exporting: {
            enabled: false
        },
        credits: {
            enabled: false
        },
        series: [{
            name: '高强度跑',
            data: []
        }, {
            name: '冲刺跑',
            data: []
        }, {
            name: '普通跑',
            data: []
        }]
    });
    return series;
}

function updateDistWithSpeedChartFirstTime(series, players) {
    var data1 = [];
    var data2 = [];
    var data3 = [];
    players.forEach(function (p, i, arr) {
        /*if (p.selected) {
         series[i].show();
         } else {
         series[i].hide();
         }*/
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.distwithspeed != null && remark.distwithspeed.length == 3) {
            data1.push(Number(remark.distwithspeed[0].toFixed(2)));
            data2.push(Number(remark.distwithspeed[1].toFixed(2)));
            data3.push(Number(remark.distwithspeed[2].toFixed(2)));
        }
    });
    series[0].setData(data1);
    series[1].setData(data2);
    series[2].setData(data3);
}

function updateDistWithSpeedChart(series, index, remark) {
    /*if (player.selected) {
     series.show();
     } else {
     series.hide();
     }*/
    var points0 = series[0].points;
    var points1 = series[1].points;
    var points2 = series[2].points;
    var ys0 = [];
    var ys1 = [];
    var ys2 = [];
    points0.forEach(function (p, i, arr) {
        ys0.push(p.y);
    });
    points1.forEach(function (p, i, arr) {
        ys1.push(p.y);
    });
    points2.forEach(function (p, i, arr) {
        ys2.push(p.y);
    });
    if (remark != null && remark.distwithspeed != null && remark.distwithspeed.length == 3 && ys0.length > index && ys1.length > index && ys2.length > index) {
        if (ys0[index] < Number(remark.distwithspeed[0].toFixed(2))) {
            ys0[index] = Number(remark.distwithspeed[0].toFixed(2));
            series[0].setData(ys0);
        }

        if (ys1[index] < Number(remark.distwithspeed[1].toFixed(2))) {
            ys1[index] = Number(remark.distwithspeed[1].toFixed(2));
            series[1].setData(ys1);
        }

        if (ys2[index] < Number(remark.distwithspeed[2].toFixed(2))) {
            ys2[index] = Number(remark.distwithspeed[2].toFixed(2));
            series[2].setData(ys2);
        }
    }
}
