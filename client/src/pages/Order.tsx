import { useState } from "react"
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"

type FoodItem = {
  id: number
  name: string
  price: number
}

const foodMenu: FoodItem[] = [
  { id: 1, name: "Burger", price: 25 },
  { id: 2, name: "Pizza", price: 40 },
  { id: 3, name: "Pasta", price: 30 },
]

export default function OrderFood() {
  const [quantities, setQuantities] = useState<{ [key: number]: number }>({})

  const updateQuantity = (id: number, value: number) => {
    setQuantities((prev) => ({
      ...prev,
      [id]: Math.max(1, value),
    }))
  }

  const addToCart = (item: FoodItem) => {
    const qty = quantities[item.id] || 1
    const total = item.price * qty
    console.log(`Added ${qty} x ${item.name} — Total: ${total} SAR`)
  }

  return (
    <div className="min-h-screen p-6 bg-gray-100 dark:bg-gray-900">
      <h1 className="text-3xl font-bold mb-6 text-center text-gray-900 dark:text-gray-100">
        Order Food
      </h1>

      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        {foodMenu.map((item) => {
          const qty = quantities[item.id] || 1
          const totalPrice = item.price * qty

          return (
            <Card key={item.id} className="p-4 bg-white dark:bg-gray-800">
              <CardHeader>
                <CardTitle className="text-gray-900 dark:text-gray-100">
                  {item.name}
                </CardTitle>
              </CardHeader>

              <CardContent className="space-y-4">
                <p className="text-lg font-semibold text-gray-900 dark:text-gray-100">
                  Price: {item.price} SAR
                </p>

                <div>
                  <label className="text-sm font-medium text-gray-900 dark:text-gray-100">
                    Quantity
                  </label>
                  <Input
                    type="number"
                    min={1}
                    value={qty}
                    onChange={(e) =>
                      updateQuantity(item.id, Number(e.target.value))
                    }
                  />
                </div>

                <p className="text-lg font-bold text-gray-900 dark:text-gray-100">
                  Total: {totalPrice} SAR
                </p>

                <Button className="w-full" onClick={() => addToCart(item)}>
                  Add to Cart
                </Button>
              </CardContent>
            </Card>
          )
        })}
      </div>
    </div>
  )
}


