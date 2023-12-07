import {
    BrowserRouter,
    Routes,
    Route,
  } from "react-router-dom";
import {Student} from './student/studentPage.js';
import {Teacher} from './teacher/teacherPage.js';
import {Home} from './home.js';
import { Principal } from "./principal/principalPage.js";

function MyRoutes(){
    return(
        <div>
            <BrowserRouter>
                <Routes>
                    <Route
                        path="/"
                        element={<Home/>}
                    />
                    <Route
                        path="/student"
                        element={<Student />}
                    />
                    <Route
                        path="/teacher"
                        element={<Teacher/>}
                    />
                    <Route
                        path="/principal"
                        element={<Principal/>}
                    />
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export {MyRoutes};