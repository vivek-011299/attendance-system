import { useState } from 'react';
import './studentPage.css';
import "bootstrap/dist/css/bootstrap.min.css";
import axios from 'axios';
import { DialogBox } from '../dialog.js';

function Student(){
    const [stu_id, set_stu_id] = useState('');
    const [student_search_box, set_student_search_box]  = useState('');
    const [state, setStateSearch] = useState(false);

    const punchin = () => {
        if(stu_id!=='')
        {
            axios.post("http://localhost:8000/student/punchin", {
                "roll":stu_id,
                "punchin":new Date().toLocaleDateString(),
                "punchout":new Date().toLocaleDateString(),
            })
        }
    }
    const punchout = () =>{

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
            <button onClick={punchin} class="btn btn-primary">Punchin</button>
            <button onClick={punchout} class="btn btn-danger">Punchout</button>
        </div>
        </>
    );
}

export {Student}