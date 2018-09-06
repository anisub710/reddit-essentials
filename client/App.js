import React from 'react';
import {Linking, StyleSheet, Text, View } from 'react-native';

export default class App extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      test: "",
    };
  }
  componentDidMount(){
    fetch('http://10.0.0.12:4000/login')
    .then(results => {
      return results.text();
      console.log(results.text());
    }).then(data => {
      this.setState({test: data});
    }).catch((error) => {
      console.error(error);
    });
  }
  render() {
    return (
      <View style={styles.container}>
        <Text>This is working: {this.state.test}</Text>
        <Text>Open up App.js to start working on your app!</Text>
        <Text>Changes you make will automatically reload.</Text>
        <Text>Shake your phone to open the developer menu.</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
