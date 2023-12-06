import './Box.css';

function Box(props) {
    return(
        <div className="box">
            <button className="box-btn">{props.text}</button>
        </div>
    );
}

export {Box};