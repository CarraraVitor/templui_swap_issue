run
```
make dev
```
- Go to http://localhost:8000
- Select any option of the SelectBox at the top of the page, which will trigger the swap of the Section with the Cards
- All the Ratings and ModalTriggers inside each Card stops working

--- 
# ORIGINAL ISSUE:

Elements Remain Uninitialized After htmxSwap

# Environment
go          go1.24.3 linux/amd64
templ       v0.3.865
templui     v0.73.2
tailwindcss v4.1.8

# Description
I hava a grid, where each grid element makes use of the Rating, Tooltip and ModalTrigger
components. 
The entire grid gets replaced through the hx-swap attribute, 
when an option of a SelectBox gets selected.
After the swap, none of the components triggers works anymore.

# Possible Solution?
After some tests and digging, I managed to figure out a way around this issue.
Here is what I believe happened:

All the event handlers for htmxSwap and htmxSwapOOB define the variable _target_
as 
```js
let target = event.detail.target || event.detail.elt;
```
This _target_ variable is then used to (re)initialize all the elements that
are related to the component in question (e.g. all the elements with the
`[data-rating-component]` attribute).
What I think is the problem is that by looking for the `detail.target` element 
first and then short circuiting with the `||` operator, the child elements that
end up being reinitialized are the ones inside the element that just got
swaped away (the target), instead of the child elements inside the new element 
being swaped into the DOM (the elt). 
Which means that all the elements removed from the DOM got
reinitialized, while the ones that just got inserted into the DOM
remains unitialized.
Defining the _target_ variable as
```js
let target = event.detail.elt || event.detail.target;
```
solved my problem.


# Demonstration

Disposition of the components in the first grid element:
- picture 

Regular behavior (without any htmxSwap):
- gif...

Wrongful(?) behavior (when target = event.detail.target || event.detail.elt):
- gif...

Fixed(?) behavior (when target = event.detail.elt || event.detail.target):
- gif...

# Conclusion (kinda of)
I didn't really know what to do with this information, but I thought it was important to share it here.
The project is not public, so I can't really share the code here. 
I am willing to share part of it privately, though.
I am also willing to open a Pull Request updating all the scripts 
if the solution presented here is found to be a valid one. Just say the word.
