import React, { Component } from 'react';
import { Alert, Text, View, ScrollView } from 'react-native';
import { Button, Dialog, Portal } from 'react-native-paper';
import { Dropdown } from 'react-native-material-dropdown';
import { WATERINGLABELS } from './GrowingSchedule/AddScheduleSelects';

//Styles
import { COMPONENTS, COLORS, FONTS, CONTAINERS } from './styles';

export class ManualWater extends Component {
  render() {
    const {
      visible,
      onDismiss,
      waterNowAmount,
      onChangeText,
      onPress,
    } = this.props;

    return (
      <Portal>
        <Dialog
          visible={visible}
          onDismiss={onDismiss}>
          <Dialog.Title>How much do you want to water? </Dialog.Title>
          <Dialog.Content>
            <Dropdown
              label={WATERINGLABELS.unitsPlaceholder}
              data={WATERINGLABELS.unitsData}
              value={waterNowAmount}
              containerStyle={COMPONENTS.dropdown}
              dropdownOffset = {{
                top: 10, 
                left: 0
              }}
              rippleOpacity={0}
              baseColor={COLORS.grey7}
              fontSize={14}
              onChangeText={onChangeText}
            />
          </Dialog.Content>
          <Dialog.Actions style={{marginRight: 10, marginBottom: 10}}>
            <Button color={COLORS.green5} onPress={onDismiss}>Cancel</Button>
            <Button color={COLORS.green5} onPress={onPress}>Go</Button>
          </Dialog.Actions>
        </Dialog>
      </Portal>
    )
  }
}

export class GenericConfirmation extends Component {
  render() {
    const {
      message,
      onDismiss,
      visible
    } = this.props;

    return (
      <Portal>
        <Dialog
          visible={visible}
          onDismiss={onDismiss}
          style={{
            padding: 10
          }}>
          <Dialog.Title>{message} </Dialog.Title>
          <Dialog.Actions style={{marginRight: 10, marginBottom: 10}}>
            <Button color={COLORS.green5} onPress={onDismiss}>Okay</Button>
          </Dialog.Actions>
        </Dialog>
      </Portal>
    )
  }
}
