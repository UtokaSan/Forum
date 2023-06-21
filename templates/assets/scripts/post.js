var url = new URL(window.location.href);
var param1Value = url.searchParams.get('param1');

fetch("/api/")