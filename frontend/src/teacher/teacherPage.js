import './teacherPage.css';
import "bootstrap/dist/css/bootstrap.min.css";

function Teacher(){
    return(
        <>
        <div className="allStudents">
                <h3>Get all Students:</h3>
                <button className='btn btn-primary'>Get all Students</button>     
        </div>
        <div className='getStudent'>
            <h3>Get student details by roll number:</h3>
            <input className='ipbox' placeholder='Enter the id here'></input>
            <button class="btn btn-info">Search</button>
        </div>
        <div className='createStudent'>
            <h3>Create student:</h3>
            <form>
                <div class="form-group">
                 <label for="usr">Name:</label>
                <input type="text" class="form-control" id="usr"/>
                </div>
                <div class="form-group">
                <label for="class">Class:</label>
                <input type="password" class="form-control" id="class"/>
                </div>
                <div class="form-group">
                <label for="age">Age:</label>
                <input type="password" class="form-control" id="age"/>
                </div>
                <div class="form-group">
                <label for="ph">Phone number:</label>
                <input type="password" class="form-control" id="ph"/>
                </div>
            </form>
            <button class="btn btn-info">Create</button>
        </div>
        </>
    );
}

export {Teacher};
