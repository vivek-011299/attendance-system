import './principalPage.css';
import "bootstrap/dist/css/bootstrap.min.css";
import { useEffect, useState } from 'react';
import axios from 'axios';
function Principal(){
    const [firstname, set_firstname] = useState('');
    const [lastname, set_lastname] = useState('');
    const [phone, set_phone] = useState('');

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
    useEffect(()=>{
        axios.get("http://localhost:8000/teacher/get_all_students")
        .then((response)=>{
            console.log(response)
        })
    })

    return(
        <>
        <div className="get_all">
            <h3>Get All students</h3>
            <button class="btn btn-primary">Get All students:</button>
        </div>
        <div className="get_id">
            <h3>Get Student with ID:</h3>
            <input placeholder="Enter the student roll number here"></input>
            <button class="btn btn-info">Search</button>
        </div>
        <div className="get_all">
            <h3>Get All teachers</h3>
            <button class="btn btn-info">Get All Teachers:</button>
        </div>
        <div className="get_id">
            <h3>Get Teacher with ID:</h3>
            <input placeholder="Enter the teacher Employee id here"></input>
            <button class="btn btn-info">Search</button>
        </div>
        <div className="create_teacher">
            <h3>Create Teacher</h3>
            <form>
                <div class="form-group">
                 <label for="first_name">First Name:</label>
                <input onChange={f_name} value={firstname} type="text" class="form-control" id="first_name"/>
                </div>
                <div class="form-group">
                <label for="last_name">Last Name:</label>
                <input onChange={l_name} value={lastname} class="form-control" id="last_name"/>
                </div>
                <div class="form-group">
                <label for="ph">Phone number:</label>
                <input onChange={phone_change} value={phone} class="form-control" id="ph"/>
                </div>
            </form>
            <button class="btn btn-info">Create</button>
            <button onClick={clear_function} class="btn btn-danger">Clear</button>
        </div>
        </>
    );
}

export {Principal};