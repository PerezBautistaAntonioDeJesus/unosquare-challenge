import React from 'react'
import { Provider } from 'react-redux'
import { store } from './combines/combine'
import { ServicesUI } from './ServicesUI'

import './styles.css'


export const UnosquareApp = () => {

  return (
    <Provider store={store}>
      <ServicesUI />
    </Provider>
  )
}
