$(document).ready(function() {

  var baseConfig = {
    highlight: true,
    minLength: 1
  };

  $('input.location').typeahead(baseConfig);
  $('input.tech').typeahead(baseConfig);


});