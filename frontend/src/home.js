import { Box } from './Box';
import './home.css';

function Home(){
    return (
        <div className='home-box'>
            <Box text="Student"/>
            <Box text="Teacher"/>
            <Box text="Principal"/>
        </div>
    );
}

export {Home};