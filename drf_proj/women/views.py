from .models import Women
from .serializers import WomenSerializer
from rest_framework.response import Response
from rest_framework.views import APIView

class WomenAPIView(APIView):
    def get(self, request):
        womens = Women.objects.all()
        return Response({'womens': WomenSerializer(womens, many=True).data})

    def post(self, requeset):
        serializer = WomenSerializer(data=requeset.data)
        serializer.is_valid(raise_exception=True)   
        serializer.save

        return Response({'post': serializer.data})

    def put(self, request, *args, **kwargs):
        pk = kwargs.get("pk", None)
        if not pk:  # Проверка на правильность запроса
            return Response({"error": "Method PUT not allowed"}, status=400)

        try:
            instance = Women.objects.get(pk=pk)  # Та запись, которую будем менять
        except:
            return Response({"error": "Object does not exist"}, status=404)
        
        # Тут мы создаём сериализатор проверяем данные и сохраняем
        serializer = WomenSerializer(instance=instance, data=request.data)
        serializer.is_valid(raise_exception=True)
        serializer.save()
        return Response(serializer.data, status=200)
    

    def delete(self, requeseet, *args, **kwargs):
        pk = kwargs.get("pk", None)
        if not pk:  # Проверка на правильность запроса
            return Response({"error": "Method PUT not allowed"}, status=400)

        Women.objects.get(pk=pk).delete()

        return Response({"post":"Delete post " + str(pk)})