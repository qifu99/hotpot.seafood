package neuralnetwork

type CacheLayer interface {
   LastInput () *SimpleMatrix
   LastOutput () *SimpleMatrix
   LastGrad () *SimpleMatrix
}

type Layer interface {
   CacheLayer
   // Output matrix shape M * N
   Dim () (int, int)
   // Calculate layer output for given input (forward propagation).
   ForwardProp (input *SimpleMatrix) *SimpleMatrix
   // Calculate input gradient.
   BackwardProp (output_grad *SimpleMatrix) *SimpleMatrix
   // Get/Set Learn Delta
   DeltaN () int
   Delta () []*SimpleMatrix
   CorrectDelta (delta []*SimpleMatrix, offset int)
   // Update layer parameter gradients as calculated from BackwardProp(). 
   ParamsUpdate (alpha float64)
}

type LossMixin interface {
   // Calculate mean loss given output and predicted output.
   Loss (output, output_pred *SimpleMatrix) *SimpleMatrix
   // Calculate input gradient given output and predicted output.
   //InputGrad (output, output_pred *SimpleMatrix) *SimpleMatrix
}

type LayerBase struct {
   Layer
   lastInput, lastOutput, lastGrad *SimpleMatrix
}

func (c *LayerBase) LastInput () *SimpleMatrix {
   return c.lastInput
}

func (c *LayerBase) LastOutput () *SimpleMatrix {
   return c.lastOutput
}

func (c *LayerBase) LastGrad () *SimpleMatrix {
   return c.lastGrad
}

func (c *LayerBase) Dim () (int, int) {
   return 0, 0
}

func (c *LayerBase) ForwardProp (input *SimpleMatrix) *SimpleMatrix {
   c.lastInput = input.Clone()
   c.lastOutput = input.Clone()
   return input
}

func (c *LayerBase) BackwardProp (output_grad *SimpleMatrix) *SimpleMatrix {
   c.lastGrad = output_grad.Clone()
   return output_grad
}

func (c *LayerBase) DeltaN () int {
   return 0
}

func (c *LayerBase) Delta () []*SimpleMatrix {
   return make([]*SimpleMatrix, 0)
}

func (c *LayerBase) CorrectDelta (delta []*SimpleMatrix, offset int) {
}

func (c *LayerBase) ParamsUpdate (alpha float64) {
}
