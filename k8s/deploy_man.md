# 如何将项目放到k8s中跑
1、在mac上电脑安装k8s
```zsh

brew install minikube
# 如果安装中途报错，清理缓存重新下载
brew cleanup kubernetes-cli
rm -rf /Users/renaxu/Library/Caches/Homebrew/downloads/*kubernetes-cli*
brew install --force kubernetes-cli
brew install minikube
```
2、 （可选）如果要用本地的镜像，构建本地镜像前一定要执行它，他的作用是将docker构建的镜像转移k8s环境中
```bash
eval $(minikube docker-env)  
```

3、将项目打包为docker镜像
进入到项目根目录,创建并编辑Dockerfile文件，执行下面的语句,
```bash
# 在这个语句前
docker build -t douyin-project-app:latest .
# 查看镜像是不是在k8s中
minikube image list | grep douyin-project-app:latest
```
4、将docker-compose.yml转换成启动k8s的.ymal文件
```bash
brew install kompose 
kompose convert -f docker-compose.yml
```
在本地项目中创建一个文件夹，用来放生成的.ymal文件，创建后进入将生成的.ymal文件放到那个位置
5、根据配置，启动集群
```bash
 kubectl apply -f .
```
6、查看是否启动成功
```bash
kubectl get pods
kubectl get deployments
```

7、😋 try try try 👀
```bash
kubectl apply -f douyin-app-deployment.yaml
kubectl describe pod douyin-app-5b7946594c-v54rp
kubectl create deployment hello-minikube --image=kicbase/echo-server:1.0
kubectl expose deployment hello-minikube --type=NodePort --port=8080
kubectl get services hello-minikube
kubectl port-forward service/hello-minikube 7080:8080\n

minikube status
minikube start
minikube dashboard
minikube stop
minikube tunnel --cleanup
minikube service hello-minikube

```