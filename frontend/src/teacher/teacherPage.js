import './teacherPage.css';
import "bootstrap/dist/css/bootstrap.min.css";
import {useState} from 'react'
import axios from 'axios';
import { DialogBox } from '../dialog.js';

function Teacher(){
    const [name, setName] = useState('');
    const [classValue,setclassValue] = useState('');
    const [age, setage] = useState('');
    const [phone, setphone] = useState('');
    const [students_data, set_students_data] = useState([]);
    const [stu_id, set_stu_id] = useState('');
    const [stu_detail, set_student_detail] = useState({});
    const [stu_attend, set_stu_attend] = useState('');
    const [stu_attend_block, set_stu_attend_block] = useState([]);
    const [created_stu, set_created_stu] = useState('');
    const [delete_stu_id,set_delete_stu_id] = useState('');
    const [delete_area, set_delete_area] = useState('');
    const [state, setStateopen] = useState(false);

    const stu_attend_box = (event) => {
        set_stu_attend(event.target.value)
    }

    const setting_roll = (event) => {
        set_stu_id(event.target.value)
    }
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
    const create_student = (e) => {
        e.preventDefault();
        if(name==='' || classValue==='' || age==='' || phone==='')
        {
            set_created_stu('Set all fields before proceeding')
        }
        else
        {
            axios.post("http://localhost:8000/teacher/create_student",{
                "name":name,
                "class":parseInt(classValue),
                "age":parseInt(age),
                "phone":phone
            })
            .then((res)=>{
                console.log(res.data)
                set_created_stu('Created')
            })
        }
    }
    const getAllStudents = () => {
        axios.get("http://localhost:8000/teacher/get_all_students")
        .then((response) => {
            set_students_data(response.data)
        })
    }
    const get_student_attend_detail = () => {
        axios.get("http://localhost:8000/teacher/get_student_attendance?id="+stu_attend)
        .then((response)=>{
            set_stu_attend_block(response.data)
        })
    }
    const delete_student_box= (e) =>{
        set_delete_stu_id(e.target.value);
    }
    const get_student_detail = () =>{
        axios.get("http://localhost:8000/teacher/get_student?id="+stu_id)
        .then((res)=>{
            set_student_detail(res.data)
        })
    }
    const delete_student = () =>{
        if(delete_stu_id==='')
        {
            set_delete_area('Please provide an id');
        }
        else{
            axios.delete("http://localhost:8000/teacher/delete_student?id="+delete_stu_id)
            .then((res)=>{
                set_delete_area('Deleted student with id '+delete_stu_id)
            })
        }
        setState();
    }
    const setState = () => {
        if(state===true)
        {
            setStateopen(false);
        }
        else{
            setStateopen(true);
        }
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
            <input className='ipbox' value={stu_id} onChange={setting_roll} placeholder='Enter the id here'></input>
            <button onClick={get_student_detail} class="btn btn-info">Search</button>
        </div>
        {
            JSON.stringify(stu_detail)!=="{}"
            &&
            <div>
                {
                    stu_detail.name
                }
            </div>
        }
        <div className='getStudentAttendance'>
            <h3>Get student attendance details by roll number:</h3>
            <input className='ipbox' value={stu_attend} onChange={stu_attend_box} placeholder='Enter the id here'></input>
            <button onClick={get_student_attend_detail} class="btn btn-info">Search</button>
        </div>
        {
            stu_attend_block.length!==0
            &&
            <div>
                {
                    stu_attend_block.map((data)=>{
                        return(
                            <div>
                                {data.roll},
                                {data.punchin},
                                {data.punchout}
                            </div>
                        )
                    })
                }
            </div>
        }
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
            <button onClick={create_student} class="btn btn-info">Create</button>
            <button onClick={handleClear} class="btn btn-danger">Clear</button>
        </div>
        {
            created_stu!==''
            &&
            <div>
                {created_stu}
            </div>
        }
        <div className='delete_box'>
            <h3>Delete a student:</h3>
            <input className='ipbox' value={delete_stu_id} onChange={delete_student_box} placeholder='Enter the id here'></input>
            <button onClick={delete_student} class="btn btn-danger">Delete</button>
        
        {   
            delete_area!==''
            &&
            <div>
                {
                <DialogBox heading="Message" message={delete_area} state={state} openFunction={setState}/>
                }
            </div>
        }
        </div>
        </>
    );
}

export {Teacher};
