import './teacherPage.css';
import "bootstrap/dist/css/bootstrap.min.css";
import {useState} from 'react'
import axios from 'axios';

function Teacher(){
    const [name, setName] = useState('');
    const [classValue,setclassValue] = useState('');
    const [age, setage] = useState('');
    const [phone, setphone] = useState('');
    const [students_data, set_students_data] = useState([]);

    const handleName = event =>{
        setName(event.target.value);
    };
    const handleClass = event =>{
        setclassValue(event.target.value);
    };
    const handleAge = event =>{
        setage(event.target.value);
    };
    const handlePhone = event => {
        setphone(event.target.value);
    };

    const handleClear = () =>{
        setName('');
        setage('');
        setclassValue('');
        setphone('');
    }
    const getAllStudents = () => {
        axios.get("http://localhost:8000/teacher/get_all_students")
        .then((response) => {
            set_students_data(response.data)
        })
    }
    return(
        <>
        <div className="allStudents">
                <h3>Get all Students:</h3>
                <button onClick={getAllStudents} className='btn btn-primary'>Get all Students</button>     
        </div>
        {
            students_data.length!==0
            &&
            <div>
                {
                    students_data.map((data) => {
                        return(
                            <div>
                                {data.name}
                            </div>
                        )
                    })
                }
            </div>
        }
        <div className='getStudent'>
            <h3>Get student details by roll number:</h3>
            <input className='ipbox' placeholder='Enter the id here'></input>
            <button class="btn btn-info">Search</button>
        </div>
        <div className='getStudentAttendance'>
            <h3>Get student attendance details by roll number:</h3>
            <input className='ipbox' placeholder='Enter the id here'></input>
            <button class="btn btn-info">Search</button>
        </div>
        <div className='createStudent'>
            <h3>Create student:</h3>
            <form>
                <div class="form-group">
                 <label for="usr">Name:</label>
                <input onChange={handleName} value={name} type="text" class="form-control" id="usr"/>
                </div>
                <div class="form-group">
                <label for="class">Class:</label>
                <input onChange={handleClass} value={classValue} class="form-control" id="class"/>
                </div>
                <div class="form-group">
                <label for="age">Age:</label>
                <input onChange={handleAge} value={age} class="form-control" id="age"/>
                </div>
                <div class="form-group">
                <label for="ph">Phone number:</label>
                <input onChange={handlePhone} value={phone} class="form-control" id="ph"/>
                </div>
            </form>
            <button class="btn btn-info">Create</button>
            <button onClick={handleClear} class="btn btn-danger">Clear</button>
        </div>
        </>
    );
}

export {Teacher};
