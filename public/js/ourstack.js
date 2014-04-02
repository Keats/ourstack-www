$(document).ready(function() {

  var baseConfig = {
    highlight: true,
    hint: false,
    minLength: 1
  };


  // We will prefetch list of techs and locations beforehand
  var getBloodhound = function(url) {
    return new Bloodhound({
      datumTokenizer: Bloodhound.tokenizers.obj.whitespace('name'),
      queryTokenizer: Bloodhound.tokenizers.whitespace,
      limit: 20,
      prefetch: {
        url: url,
        filter: function(list) {
          return $.map(list, function(item) { return { name: item }; });
        }
      }
    });
  }

  var techs = getBloodhound('/techs');
  var locations = getBloodhound('/locations');

  techs.initialize();
  locations.initialize();

  $('input.location').typeahead(baseConfig, {
    name: 'locations',
    displayKey: 'name',
    source: locations.ttAdapter()
  });

  $('input.tech').typeahead(baseConfig, {
    name: 'techs',
    displayKey: 'name',
    source: techs.ttAdapter()
  });

});
