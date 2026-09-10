import { useState, useEffect } from 'react' ;
import './App.css' 
function App() { 
  const [price, setPrice] = useState(null)
   useEffect(() => { async function fetchPrice() {
     try { const response = await fetch( `https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=IBM&apikey=${import.meta.env.VITE_ALPHA_VANTAGE_API}` )
      if (!response.ok) {
         throw new Error('Network response was not ok') } 
         const result = await response.json() 
         setPrice(result) 
        } catch (error) { 
          console.error('Error fetching stock data:', error) 
        }
       }
         fetchPrice() }, [])
          return ( <div> <h1>IBM Stock Data</h1> 
          {price ? ( <pre>{JSON.stringify(price, null, 2)}</pre> ) : ( <p>Loading...</p> )} 
          </div> ) }
           export default App