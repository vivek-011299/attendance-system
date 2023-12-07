import { useNavigate } from 'react-router-dom';
import './Box.css';

function Box(props) {
    const navigate = useNavigate();
    const NavigateStudent = () => {
        if (props.text==="Student")
            navigate("/student");
        if(props.text==="Teacher")
            navigate("/teacher");
        if(props.text==="Principal")
            navigate("/principal");
    }
    return(
        <div className="box">
            <button onClick={NavigateStudent} className="box-btn">{props.text}</button>
        </div>
    );
}

export {Box};