import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { activeService } from '../actions/services';

export const GridItem = (service) => {
  const dispatch = useDispatch();
  const  { active } = useSelector( state => state.services)

  console.log( active,  )

  const { metadata} = service;

  const handleOnclick = () => {
    dispatch( activeService(service))
  }
  return (
    <>
      <p className={`serviceName ${(active !== null && active.metadata.uid === metadata.uid) ? 'active' : ''}`} onClick={ handleOnclick }> {metadata.name} </p>
    </>
  )
}
