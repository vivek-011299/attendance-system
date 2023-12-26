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
    const [state_punch, set_state_punch] = useState(false);
    const [head,set_head] = useState('');


    const punch_state = () =>{
        if(state_punch===true)
        {
            set_state_punch(false);
        }
        else
            set_state_punch(true);
    }
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
                set_msg(res.data['message'])
                if(res.data['message']==="Punched in")
                {
                    set_head("Success!")
                }
                else{
                    set_head('Error');
                }
            })
        }
        set_state_punch(true)
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
                set_msg(res.data['message'])
                if(res.data['message']==='Punched out')
                {
                    set_head("Success!")
                }
                else{
                    set_head('Error');
                }
            })
        }
        set_state_punch(true)
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
                console.log(res.data)
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
                    <DialogBox heading="Message" message={student_search_box} state={state} openFunction={setState}/>
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
            <div>
                <DialogBox heading={head} message={msg} openFunction={punch_state} state={state_punch} />
            </div>
        }
        </>
    );
}

export {Student}