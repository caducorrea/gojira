$(document).ready(function(){
var ticketLink = document.getElementById('ticketLink');
ticketLink.onclick = function(){
	ReactDOM.render(
  <TicketSearch />,
  document.getElementById('content')
);



}});