import React from 'react'

import { useStateContext } from '../../contexts/ContextProvider';

const LineChart = () => {
  const { currentMode } = useStateContext();

  return (
    <div className=''>
      Line chart
    </div>
  );
};

export default LineChart;