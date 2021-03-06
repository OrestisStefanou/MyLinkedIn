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
import ProfessionalProfile from "./profile/ProfessionalProfile";
import MessagesPage from "./homepage/MessagesPage";
import ChatPage from "./homepage/ChatPage";
import JobAdsPage from "./homepage/JobAdsPage";
import UserJobAdsPage from "./homepage/UserJobAdsPage";
import ErrorPage from "./errorPage";
import AdminSignIn from "./admin/AdminSignIn";
import AdminHomepage from "./admin/AdminHomepage";
import UsersPage from "./admin/UsersPage";
import UserProfile from "./admin/UserProfile";

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
                <Route exact path='/settings'>
                    <Profile/>
                </Route>
                <Route exact path='/notifications'>
                    <NotificationsPage/>
                </Route>
                <Route exact path='/network'>
                    <NetworkPage/>
                </Route>
                <Route exact path='/messages'>
                    <MessagesPage/>
                </Route>
                <Route exact path='/jobs'>
                    <JobAdsPage/>
                </Route>
                <Route exact path='/myAds'>
                    <UserJobAdsPage/>
                </Route>
                <Route exact path='/professionals/:email' children={<ProfessionalProfile/>}></Route>
                <Route exact path='/chat/:id' children={<ChatPage/>}></Route>
                {/*Admin routes */}
                <Route exact path='/admin/signin'>
                    <AdminSignIn/>
                </Route>
                <Route exact path='/admin/home'>
                    <AdminHomepage/>
                </Route>
                <Route exact path='/admin/users'>
                    <UsersPage/>
                </Route>
                <Route exact path='/admin/user/:email' children={<UserProfile/>}></Route>
                <Route path='*'>
                    <ErrorPage />
                </Route>
            </Switch>
        </Router>
    )
}

ReactDom.render(<App/>,document.getElementById('root'));
