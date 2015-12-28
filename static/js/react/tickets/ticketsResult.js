window.TicketResult = React.createClass({
	render: function(){

				console.log("started at: " + new Date());
				var rows = [];
				this.props.tickets.forEach(function(ticket, index){
					rows.push(<tr>
	      							<td className="mdl-data-table__cell--non-numeric">{ticket.ID}</td>
	      							<td className="mdl-data-table__cell--non-numeric">{ticket.Title}</td>
	      							<td className="mdl-data-table__cell--non-numeric">{ticket.Department}</td>
	      							<td className="mdl-data-table__cell--non-numeric">{ticket.Description}</td>
	    						 </tr>);
				});

		return (
			<table className="mdl-data-table mdl-js-data-table">
  <thead>
    <tr>
      <th className="mdl-data-table__cell--non-numeric">ID</th>
      <th className="mdl-data-table__cell--non-numeric">Title</th>
      <th className="mdl-data-table__cell--non-numeric">Department</th>
      <th className="mdl-data-table__cell--non-numeric">Description</th>
    </tr>
  </thead>
  <tbody>
   {rows}
  </tbody>
</table>
			);
	}
	
});