import React, {Component} from 'react';
import _ from 'lodash';
import { Button, Card, Divider } from 'react-native-paper';
import { StyleSheet, Modal, Text, TouchableHighlight, View, ScrollView } from 'react-native';
import Icon from 'react-native-vector-icons/Ionicons';
import { default as AwesomeIcon } from 'react-native-vector-icons/FontAwesome';

import { COLORS, CONTAINERS, COMPONENTS, FONTS } from '../styles';
import AddScheduleForm from './AddScheduleForm';

export default class ListViewCard extends Component {
  state = {
    eventListEnd: 6, //initially the list has 3 elements, the array is from 0 to 2, then add the dividers
    viewable: true, //remove View More button if end of eventList is reached
    modalVisible: false,
    count: 0
  }

  setModalVisible = (visible) => {
    this.setState({modalVisible: visible});
  }


  viewMore = (eventList) => {
    if(this.state.eventListEnd + 6 > eventList.length ){
      this.setState({
        eventListEnd: eventList.length,
        viewable: false
      })
    } else {
      this.setState({
        eventListEnd: this.state.eventListEnd + 6
      })
    }
  }

  render() {
    let iconComponent;
    let iconName=this.props.iconName;

    if(iconName === 'ios-water') {
      iconComponent = <Icon name={iconName} size={100} color={COLORS.green5}/>;
    } else if(iconName === 'lightbulb-o') {
      iconComponent = <AwesomeIcon name={iconName} size={100} color={COLORS.green5}/>;
    }

    let eventList = [];
    _.forEach(this.props.items, event => {
      eventList.push(<Divider />); //push separately because elements need parents
      eventList.push(      
        <View style={{ //not sure why it won't let me add bottom padding without stlying like this
          flex: 1, 
          flexDirection: 'row', 
          justifyContent: 'space-between',
          paddingTop: 10, 
          paddingBottom: 10
        }}
        key={event.id}
        >
          <Text style={{fontWeight: 'bold'}}>{event.date}</Text>     
          <Text>{event.time}</Text>
          <Text>{event.amount}</Text>
        </View>
        );
    });

    if(eventList.length === 0) {
      eventList.push(
        <Text style={{textAlign: 'center', paddingBottom: 15}}>No events to show.</Text>
      );
    }

    return (  
      <Card style={CONTAINERS.listViewCard}>
        <Modal
          animationType="slide"
          transparent={false}
          visible={this.state.modalVisible}
          onRequestClose={() => {
            Alert.alert('Modal has been closed.');
          }}>
            <ScrollView>
              <AddScheduleForm
                hide={this.setModalVisible}
                parent={this.props.iconName}
                schedules={this.props.fullSchedule}
                addSchedule={this.props.addSchedule}
                deleteSchedule={this.props.deleteSchedule}
              />
            </ScrollView>

          </Modal>
        <Card.Content>
          <View style={_.assignIn(CONTAINERS.spaceBetween, {paddingTop: '12%'})}> 
            {iconComponent}

            <View style={{width: '70%', paddingTop: '15%'}}>
              <Text style={_.assignIn(FONTS.h3, {textAlign: 'center'})}>
                {this.props.message}
              </Text>
            </View>
          </View>

          <Text style={{paddingBottom: 10}}>Upcoming Events</Text>
          { _.slice(eventList, 0, this.state.eventListEnd)}

          {this.state.viewable && 
          <View 
            style={{        
              alignItems: "center", 
              paddingTop: 15,
              paddingBottom: 15}}>
            <Button 
              mode='text'
              color={COLORS.grey5}
              onPress={() => this.viewMore(eventList)}>
              View More
            </Button>
            </View> }

          <View style={_.assignIn(CONTAINERS.spaceBetween, {paddingTop: '25%'})}> 
            <Button 
              mode='contained'
              color={COLORS.green5}
              onPress={this.props.mainButtonFunction}>
              {this.props.mainButtonName /*water now or turn light on*/}
            </Button>
            <Text               
              onPress={() => {
                this.setModalVisible(true);
              }}
              style={{
                color: COLORS.green5,
                width: 150,
                textTransform: 'uppercase',
                paddingLeft: 25,
                fontWeight: 'bold'
              }}
            >
              Add/Remove Schedules
            </Text>
          </View>
        </Card.Content>
      </Card>
    );
  }
}
