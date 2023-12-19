import React from "react";
import Dialog from "@material-ui/core/Dialog";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import Button from "@material-ui/core/Button";

function DialogBox(props){
    return(
        <div stlye={{}}>
			<Dialog open={props.state} onClose={props.openFunction}>
				<DialogTitle>{props.heading}</DialogTitle>
				<DialogContent>
					<DialogContentText>
						{props.message}
					</DialogContentText>
				</DialogContent>
				<DialogActions>
					<Button onClick={props.openFunction}
						color="primary" autoFocus>
						Close
					</Button>
				</DialogActions>
			</Dialog>
		</div>
    );
}

export {DialogBox};


