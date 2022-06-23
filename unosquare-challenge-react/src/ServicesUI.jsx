import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';
import { startServices } from './actions/services';

import { GridItem } from './components';
import { SelectedItem } from './components/SelectedItem';

export const ServicesUI = () => {

    const dispatch = useDispatch();

    useEffect(() => {
        dispatch( startServices() );
    }, [ dispatch ])

    const { services, active} =  useSelector( state => state.services )
    

  return (
    <>
    
    <h1 className='titlePage'>Services</h1>
    <div className='container'>

        <div className='sidebar'>
            {
                services.map( service => (
                    <GridItem key={service.metadata.uid} {...service} />
                ))
            }

        </div>

        <div className='selectedItem'>
            {
                active && <SelectedItem />
            }
        </div>
    </div>
    
    </>
  )
}
