# task-operator


# development
operator-sdk new task-operator --repo=github.com/umutkocasarac/task-operator

operator-sdk add api --api-version=umut.com/v1alpha1 --kind=Task

Update the task type spec

	Image string `json:"image"`
	Factorial int `json:"factorial"`
Add following to status
	Result string `json:"result"`

operator-sdk generate k8s
operator-sdk generate crds

operator-sdk add controller --api-version=umut.com/v1alpha1 --kind=Task

Update the file under crds/deploy
apiVersion: umut.com/v1alpha1
kind: Task
metadata:
  name: example-task
spec:
  # Add fields here
  image: ukocasarac/test:1
  factorial: 4
status:
  result: "not-created"

# deployment

**Create the service account etc on the kubernetes**
kubectl apply -f deploy/service_account.yaml
kubectl apply -f deploy/role.yaml
kubectl apply -f deploy/role_binding.yaml
kubectl apply -f deploy/operator.yaml
kubectl apply -f deploy/crds/umut.com_tasks_crd.yaml

**Run the job locally**
export OPERATOR_NAME=task-operator
operator-sdk run --local

**Deploy new task to kubernetes, it will create a job with pod**
kubectl apply -f deploy/crds/umut.com_v1alpha1_task_cr.yaml

**check the pods and jobs with following command**
kubectl get jobs
kubectl get pods
kubectl describe task example-task