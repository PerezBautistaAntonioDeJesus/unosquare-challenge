import React from 'react'
import { useSelector } from 'react-redux'

export const SelectedItem = () => {

    const { active } = useSelector( state => state.services);
    console.log( active )
    const { metadata, spec } = active;
  return (
    <>
    <div>{ metadata.name }</div>

    <div className='selectedItemContent'>
      <p className='attribute'><em>Name:</em> {metadata.name} </p>
      <p className='attribute'><em>Namespace:</em> {metadata.namespace} </p>
      <p className='attribute'><em>Resource Version:</em> {metadata.resource_version} </p>
      <p className='attribute'><em>UID:</em> {metadata.uid} </p>
    </div>
    <div className='selectedItemContent'>
      <p className='attribute'><em>ClusterIP:</em> {spec.cluster_ip} </p>
      <p className='attribute'><em>ClustersIPs:</em> {spec.cluster_ips} </p>
      <p className='attribute'><em>Iternal Traffic Policy:</em> {spec.internal_traffic_policy} </p>
      <p className='attribute'><em>IP Families:</em> {spec.ip_families} </p>
      <p className='attribute'><em>Session Affinity:</em> {spec.session_affinity} </p>
      <p className='attribute'><em>Type:</em> {spec.type} </p>
    </div>
    </>
  )
}