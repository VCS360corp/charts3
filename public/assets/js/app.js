// var barColors      = [ '#008B8B', '#66CCCC', '#61B329', '#BCED91', '#9932CC', '#7D26CD'],
//     highlightColor = '#FFF68F',
// legendLabels  = ['Cats','Cats playing music instruments','Awesome Code Examples','Narwhal sightings','Grahing libraries','Kittens'],
// data           = [45, 30, 10, 5, 5, 5];

function drawChart(el, data) {
  var labels = [];
  var count = [];
  var colors = [];
  console.log(data);
  
 data.forEach(function(item){
    labels.push(item.category);
    count.push(item.count);
    colors.push('asdas')
  })
  console.log(count)
  console.log(colors)
  console.log(labels)
  var $container = $(el);
   $container.html('');

  var barColors = ['#008B8B'],
    highlightColor = '#FFF68F',
    legendLabels = labels,
    data = count;

  var pHeight = parseInt($container.css('height')),
    pWidth = parseInt($container.css('width')),
    radius = pWidth < pHeight ? pWidth / 3 : pHeight / 3;
  bgColor = jQuery('body').css('background-color');

  var paper = new Raphael($container[0], pWidth, pHeight);

  // draw the piechart
  var pie = paper.piechart(pWidth / 2, pHeight / 2, radius, data, {
    legend: legendLabels,
    legendpos: 'east',
    legendcolor: '#eaeaea',
    stroke: bgColor,
    strokewidth: 1,
    // colors: barColors
  });

  // assign the hover in/out functions
  pie.hover(function () {
    this.sector.stop();
    this.sector.scale(1.1, 1.1, this.cx, this.cy);
    this.sector.animate({ 'stroke': highlightColor }, 400);
    this.sector.animate({ 'stroke-width': 1 }, 500, "bounce");

    if (this.label) {
      this.label[0].stop();
      this.label[0].attr({ r: 8.5 });
      this.label[1].attr({ "font-weight": 800 });
      center_label.attr('text', this.value.value + '%');
      center_label.animate({ 'opacity': 1.0 }, 200);
    }
  }, function () {
    this.sector.animate({ transform: 's1 1 ' + this.cx + ' ' + this.cy }, 500, "bounce");
    this.sector.animate({ 'stroke': bgColor }, 400);
    if (this.label) {
      this.label[0].animate({ r: 5 }, 500, "bounce");
      this.label[1].attr({ "font-weight": 400 });
      //center_label.attr('text','');
      center_label.animate({ 'opacity': 0.0 }, 500);
    }
  });

  // blank circle in center to create donut hole effect
  paper.circle(pWidth / 2, pHeight / 2, radius * 0.6)
    .attr({ 'fill': bgColor, 'stroke': bgColor });

  var center_label = paper.text(pWidth / 2, pHeight / 2, '')
    .attr({ 'fill': '#eaeaea', 'font-size': '18', "font-weight": 800, 'opacity': 0.0 });
}


$(document).ready(function () {
  loadCats();
})

function loadDogs(){
  $('#dogs').attr('disabled', true);
  $('#cats').attr('disabled', false);
  $.get(window.app.url + "/dogs", function (data) {
    drawChart('#graph', data);
  })
}
function loadCats(){
  $('#cats').attr('disabled', true);
  $('#dogs').attr('disabled', false);
  $.get(window.app.url + "/cats", function (data) {
    drawChart('#graph', data);
  })
}