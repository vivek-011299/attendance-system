import './principalPage.css';
import "bootstrap/dist/css/bootstrap.min.css";
import { useState } from 'react';
import axios from 'axios';
import { DialogBox } from '../dialog.js';

function Principal(){
    const [firstname, set_firstname] = useState('');
    const [lastname, set_lastname] = useState('');
    const [phone, set_phone] = useState('');
    const [roll, setroll] = useState('');
    const [emp_id, set_emp_id] = useState('');
    const [all_data, set_all_data] = useState([]);
    const [data_with_id, set_data_with_id] = useState({});
    const [teacher_data, set_teacher_data] = useState([]);
    const [teacher_id_data, set_teacher_id_data] = useState({});
    const [delete_area, set_delete_area] = useState('');
    const [delete_t_id, set_delete_t_id] = useState('');
    const [state, setStateforDelete] = useState(false);

    const setState = () =>{
        if(state===true)
        {
            setStateforDelete(false);
        }
        else{
            setStateforDelete(true);
        }
    }

    const delete_teacher_box = (e) =>{
        set_delete_t_id(e.target.value);
    }

    const f_name = event =>{
        set_firstname(event.target.value);
    };
    const l_name = event => {
        set_lastname(event.target.value)
    };
    const phone_change = event => {
        set_phone(event.target.value);
    };
    const clear_function = () =>{
        set_firstname('');
        set_lastname('');
        set_phone('');
    };

    const getAllStudents = () => {
        axios.get("http://localhost:8000/principal/get_all_students")
        .then((response)=>{
            set_all_data(response.data);
            console.log(response);
        })
    }
    const setRoll = (event) => {
        setroll(event.target.value);
    }
    const getStudentwithID = () => {
        axios.get(`http://localhost:8000/principal/get_student?id=${roll}`)
        .then((res) => {
            set_data_with_id(res.data)
        })
    }
    const teacher_input = (event) => {
        set_emp_id(event.target.value);
    }
    const getAllTeachers = () => {
        axios.get("http://localhost:8000/principal/get_all_teachers")
        .then((res) => {
            set_teacher_data(res.data);
        })
    }
    const teacher_with_id = () => {
        axios.get(`http://localhost:8000/principal/get_teacher?id=${emp_id}`)
        .then((response)=>{
            set_teacher_id_data(response.data)
        })
    }
    const create_teacher = (e) =>{
        e.preventDefault();
        axios.post("http://localhost:8000/principal/create_teacher",{
            "firstname":firstname,
            "lastname":lastname,
            "phonenumber":phone
        })
        .then((res)=>{
            console.log(res)
        })
    }
    const delete_teacher = () => {
        if(delete_t_id!=='')
        {
            axios.delete("http://localhost:8000/principal/delete_teacher?id="+delete_t_id)
            .then((res)=>{
                set_delete_area('Deleted successfully')
            })
        }
        else{
            set_delete_area("Please provide an id to delete");
        }
        setState();
    }
    return(
        <>
        <div className="get_all">
            <h3>Get All students</h3>
            <button class="btn btn-primary" onClick={getAllStudents}>Get All students:</button>
        </div>
        {all_data.length !==0
         && 
        <div>
        {
                all_data.map((data)=>{
                    return (
                        <div>
                            {data.roll}
                        </div>
                    );

                })
            }
        </div>
}
        <div className="get_id">
            <h3>Get Student with ID:</h3>
            <input value={roll} onChange={setRoll} placeholder="Enter the student roll number here"></input>
            <button onClick={getStudentwithID} class="btn btn-info">Search</button>
        </div>
        {
            JSON.stringify(data_with_id) !== "{}"
            &&
            <div>
            {
                <div>
                    {data_with_id.name}
                </div>
            }
            </div>
        }
        
        <div className="get_all">
            <h3>Get All teachers</h3>
            <button onClick={getAllTeachers} class="btn btn-info">Get All Teachers:</button>
        </div>
        {
            teacher_data.length!==0 &&
            <div>
                {
                    teacher_data.map((data) => {
                        return (
                        <div>
                            {data.firstname}
                        </div>
                    )
                })
                }
            </div>
        }
        <div className="get_id">
            <h3>Get Teacher with ID:</h3>
            <input value={emp_id} onChange={teacher_input} placeholder="Enter the teacher Employee id here"></input>
            <button onClick={teacher_with_id} class="btn btn-info">Search</button>
        </div>
        {
            JSON.stringify(teacher_id_data) !== "{}"
            &&
            <div>
                {
                    teacher_id_data.firstname
                }
            </div>
        }
        <div className="create_teacher">
            <h3>Create Teacher</h3>
            <form>
                <div class="form-group">
                 <label for="first_name">First Name:</label>
                <input onChange={f_name} value={firstname} type="text" class="form-control" id="first_name"/>
                </div>
                <div class="form-group">
                <label for="last_name">Last Name:</label>
                <input onChange={l_name} value={lastname} type='text' class="form-control" id="last_name"/>
                </div>
                <div class="form-group">
                <label for="ph">Phone number:</label>
                <input onChange={phone_change} value={phone} type='text' class="form-control" id="ph"/>
                </div>
            </form>
            <button onClick={create_teacher} class="btn btn-info">Create</button>
            <button onClick={clear_function} class="btn btn-danger">Clear</button>
        </div>
        <div className='delete_box'>
            <h3>Delete a teacher:</h3>
            <input className='ipbox' value={delete_t_id} onChange={delete_teacher_box} placeholder='Enter the id here'></input>
            <button onClick={delete_teacher} class="btn btn-danger">Delete</button>
        
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

export {Principal};