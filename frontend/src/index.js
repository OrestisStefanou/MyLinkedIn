import React from 'react';
import ReactDom from 'react-dom';
import {BrowserRouter as Router,Route,Switch} from 'react-router-dom';
import SignUp from "./SignUp";
import SignIn from "./SignIn";
import WeclomePage from "./WelcomePage";
import Homepage from './homepage/Homepage';
import Profile from "./profile/Profile"
import NotificationsPage from "./homepage/NotificationsPage"
import NetworkPage from "./homepage/NetworkPage";
import ProfessionalProfile from "./profile/ProfessionalProfile"

const App = () => {
    return(
        <Router>
            <Switch>
                <Route exact path='/'>
                    <WeclomePage/>
                </Route>
                <Route exact path='/signup'>
                    <SignUp/>
                </Route>
                <Route exact path='/signin'>
                    <SignIn/>
                </Route>
                <Route exact path='/home'>
                    <Homepage/>
                </Route>
                <Route exact path='/profile'>
                    <Profile/>
                </Route>
                <Route exact path='/notifications'>
                    <NotificationsPage/>
                </Route>
                <Route exact path='/network'>
                    <NetworkPage/>
                </Route>
                <Route exact path='/professionals/:email' children={<ProfessionalProfile/>}></Route>
            </Switch>
        </Router>
    )
}

ReactDom.render(<App/>,document.getElementById('root'));
