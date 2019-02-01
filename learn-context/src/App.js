import React, { Component } from 'react';
import * as Unstated from 'unstated';

import logo from './logo.svg';
import './App.css';

const ThemeContext = React.createContext('defaultTheme')
const UserContext = React.createContext('defaultUser')

const Store = React.createContext({
  user: {
    name: 'Default User'
  },
  settings: {
    theme: 'dark'
  }
});

class PersistedContainer extends Unstated.Container {
  constructor(...args) {
    super(...args)
  } 

  async setState(...newState) {
    super.setState(...newState).then(() => {
      console.log('updated state', this.state)
    })
  }
}

class UserStore extends PersistedContainer {
  state = {
    firstName: 'First',
    lastName: 'Last'
  }

  changeFirstName = (firstName = '') => {
    this.setState({ firstName })
  }

  changeLastName = (lastName = '') => {
    this.setState({ lastName })
  }

  getName = () => `${this.state.firstName} ${this.state.lastName}`
}

class ThemeStore extends Unstated.Container {
  state = {
    theme: 'dark'
  }

  getTheme = () => `${this.state.theme}`
}

class ContextApp extends Component {
  render() {
    return (
      <React.Fragment>
        <React.Fragment>
          {/* separate */}
          <ThemeContext.Consumer>
            {theme => (
              <p>{theme}</p>
            )}
          </ThemeContext.Consumer>
        </React.Fragment>
        {/* separate */}
        <React.Fragment>
          <UserContext.Consumer>
            {user => (
              <p>{user}</p>
            )}
          </UserContext.Consumer>
        </React.Fragment>
        {/* nested consumer */}
        <React.Fragment>
          <ThemeContext.Consumer>
            {theme => (
              <UserContext.Consumer>
                {user => (
                  <p>{theme} {user}</p>
                )}
              </UserContext.Consumer>
            )}
          </ThemeContext.Consumer>
        </React.Fragment>
        {/* store test 1 */}
        <React.Fragment>
          <Store.Consumer>
            {({ user, settings }) => (
              <p>{user.name}, {settings.theme}</p>
            )}
          </Store.Consumer>
        </React.Fragment>
        {/* unstated test */}
        <Unstated.Provider>
          <Unstated.Subscribe to={[UserStore, ThemeStore]}>
            {(userStore, themeStore) => (
              <React.Fragment>
                <p>{userStore.getName()}</p>
                <p>{themeStore.getTheme()}</p>
                <button onClick={() => userStore.changeFirstName(Date.now())}>Change first name</button>
                <button onClick={() => userStore.changeLastName(Date.now())}>Change last name</button>
              </React.Fragment>
            )}
          </Unstated.Subscribe>
        </Unstated.Provider>
      </React.Fragment>
    )
  }
}

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        <div>
          <ContextApp />
        </div>
      </div>
    );
  }
}

export default App;
