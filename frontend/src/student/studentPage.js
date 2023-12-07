import './studentPage.css';
import "bootstrap/dist/css/bootstrap.min.css";

function Student(){
    return(
        <>
        <div className='inputbox'>
            <h2>Enter your ID in the input box</h2>
            <input className="searchbox" placeholder='Enter your id here'></input>
            <button class="btn btn-info">Search</button>
        </div>
        <div className='pipout-buttons'>
            <button class="btn btn-primary">Punchin</button>
            <button class="btn btn-danger">Punchout</button>
        </div>
        </>
    );
}

export {Student}