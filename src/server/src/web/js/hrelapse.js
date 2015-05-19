/**
 * Created by piasy on 15/5/18.
 */

var NAMES = ['~ 100', '100 ~ 120', '120 ~ 140', '140 ~ 160', '160 ~ 180', '180 ~ 190', '190 ~ 200', '200 ~'];

function createHRElapseChart(players) {
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });

    var series = [];
    $('#chart_container').highcharts({
        chart: {
            type: 'pie',
            options3d: {
                enabled: true,
                alpha: 45,
                beta: 0
            },
            events: {
                load: function() {
                    series = this.series;
                }
            }
        },
        title: {
            text: '不同心率持续时间比例'
        },
        tooltip: {
            pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                depth: 35,
                dataLabels: {
                    enabled: true,
                    format: '{point.name}'
                }
            }
        },
        series: [{
            type: 'pie',
            data: []
        }],
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
        }
    });
    return series;
}

function updateHRElapseChartFirstTime(series, players) {
    var data = [];
    players.forEach(function (p, i, arr) {
        /*if (p.selected) {
         series[i].show();
         } else {
         series[i].hide();
         }*/
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.hrelapse != null && remark.hrelapse.length == 8) {
            var sum = 0;
            remark.hrelapse.forEach(function (elapse, j, arr) {
                sum += elapse;
            });
            if (sum > 0) {
                remark.hrelapse.forEach(function (elapse, j, arr) {
                    data.push([NAMES[j], elapse / sum]);
                });
            }
        }
    });
    series[0].setData(data);
}

