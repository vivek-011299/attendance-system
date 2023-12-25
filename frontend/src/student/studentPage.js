import { useState } from 'react';
import './studentPage.css';
import "bootstrap/dist/css/bootstrap.min.css";
import axios from 'axios';
import { DialogBox } from '../dialog.js';

function Student(){
    const [stu_id, set_stu_id] = useState('');
    const [student_search_box, set_student_search_box]  = useState('');
    const [state, setStateSearch] = useState(false);
    const [msg, set_msg] = useState('');

    const punchin = () => {
        if(stu_id!=='')
        {
            const date = new Date();

            axios.post("http://localhost:8000/student/punchin", {
                "roll": parseInt(stu_id),
                "punchin":date.getHours()+":"+date.getMinutes()+":"+date.getSeconds(),
                "punchout":'0',
            })
            .then((res)=>{
                set_msg('Student punched in')
            })
        }
    }
    const punchout = () =>{
        if(stu_id!=='')
        {
            const date = new Date();
            axios.post("http://localhost:8000/student/punchout",{
                "roll":parseInt(stu_id),
                "punchout":date.getHours()+":"+date.getMinutes()+":"+date.getSeconds()
            })
            .then((res)=>{
                set_msg('Student punched out')
            })
        }
    }

    const stu_id_change = (e) => {
        set_stu_id(e.target.value);
    }

    const setState = () =>{
        if(state===true)
        {
            setStateSearch(false);
        }
        else{
            setStateSearch(true);
        }
    }

    const searchStudent = () => {
        if(stu_id !=='')
        {
            axios.get("http://localhost:8000/student/search?id="+stu_id)
            .then((res)=>{
                if(res.data['roll']!==0)
                {
                    set_student_search_box('Student with id '+stu_id+' is present')
                }
                else{ 
                    set_student_search_box('Student with id '+stu_id+' is NOT present')
                }
            })
        }
        else{
            set_student_search_box("Please provide an id");
        }
        setState()
    }
    return(
        <>
        <div className='inputbox'>
            <h2>Enter your ID in the input box</h2>
            <input className="searchbox" value={stu_id} onChange={stu_id_change} placeholder='Enter your id here'></input>
            <button onClick={searchStudent} class="btn btn-info">Search</button>
        </div>
        {
            <div className='SearchText'>
                {
                    <DialogBox heading="message" message={student_search_box} state={state} openFunction={setState}/>
                }
            </div>
        }
        <div className='pipout-buttons'>
            {
                student_search_box.includes('is present')===true
                &&
                <div>
                <button onClick={punchin} class="btn btn-primary">Punchin</button>
                <button onClick={punchout} class="btn btn-danger">Punchout</button>
                </div>
            }
        </div>
        {
            msg!==''
            &&
            <div className='msg_box'>
                <h4>{msg}</h4>
            </div>
        }
        </>
    );
}

export {Student}