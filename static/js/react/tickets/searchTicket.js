window.TicketSearch = React.createClass({
		
handleClick: function(e){
		var searchTicket = document.getElementById('text_search').value;
		$.ajax({
			url: "/tickets?search_term="+searchTicket,
			
			method:"GET"})
		.done(function(data){
				ReactDOM.render(<TicketResult tickets={data.tickets} />, document.getElementById('ticketResult'));
			})
		.fail(function(data)
			{
				toastr.error('Error on search!');
			});
			
		},
	render: function(){
		return(
			<div>
		<div className="mdl-grid">
  <div className="mdl-typography--headline mdl-cell mdl-cell--12-col">Search</div>
  <div className="mdl-textfield mdl-js-textfield mdl-cell mdl-cell--10-col">
    <input className="mdl-textfield__input" type="text" id="text_search" />
    <label className="mdl-textfield__label" forHtml="text_search">Type something to search your tickets</label>
  </div>
</div>
<button id="searchTicket" className="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect"  
onClick={this.handleClick}>Search</button>
  

<div id="orderBusyIndicator" className="mdl-spinner mdl-js-spinner"></div>
<div id="ticketResult"></div>
</div>
			);
	}
});

